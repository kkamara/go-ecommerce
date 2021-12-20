@extends('layouts.master')

@section('content')
    <div class="row">
        <div class="col-md-9">
            <table class="table table-striped">
                <tbody>
                    <tr class='text-center'>
                        <th scope='row' colspan='2'>
                            <img style='max-height:100px' src="{{ $product->image_path }}" class='img-responsive'/>
                            <h4 dusk="product-name">
                                {{ $product->name }}
                            </h4>
                        </th>
                    </tr>
                    <tr>
                        <td scope='row'>Description</td>
                        <td>
                            {{ $product->short_description }}
                            <br/>
                            <br/>
                            {{ $product->long_description }}
                        </td>
                    </tr>
                    <tr>
                        <td scope='row'>Product Details</td>
                        <td>
                            {{ $product->product_details }}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div class="col-md-3">
            <ul class="list-group">
                <li class="list-group-item">
                    <h3>
                        {{ $product->formatted_cost }}
                    </h3>
                    <small>
                        @if($product->shippable)
                            This item is shippable
                        @else
                            This item is <strong>not</strong> shippable
                        @endif
                    </small>
                    <br/>
                    @if($product->free_delivery)
                        <small>Free Delivery</small>
                    @endif
                </li>
                <li class="list-group-item">
                    @if(Auth::check() && $product->doesUserOwnProduct())
                        <a 
                            dusk="edit-btn"
                            href='{{ route('companyProductEdit', [$product->company->slug, $product->id]) }}' 
                            class='btn btn-warning btn-sm pull-left'
                        >
                            Edit item
                        </a>
                        <a 
                            dusk="delete-btn"
                            href='{{ route('companyProductDelete', [$product->company->slug, $product->id]) }}' 
                            class='btn btn-danger btn-sm pull-right'
                        >
                            Delete item
                        </a>
                    @else
                        <a 
                            dusk="add-to-cart-btn"
                            href='{{ route('productAdd', $product->id) }}' 
                            class='btn btn-primary btn-sm'>
                            Add to cart
                        </a>
                    @endif
                </li>
            </ul>
        </div>
    </div>
    <div class="row">
        <div class="col-md-12">
            <div class="card">
                <div class="card-header">
                    <div class="lead">
                        Reviews @if($product->review !== '0.00') (Average  {{ $product->review }}) @endif
                    </div>
                </div>
                <div class="card-body">
                    <div class="card-text">

                        @forelse($reviews as $review)
                            <div class="card">
                                @if(! $review->isFlaggedFiveTimes())
                                <div class="card-body">
                                    <div class="card-text">
                                        <div class="float-left">
                                            Product Rated {{ $review->score }} / 5

                                            @if(Auth::check() && auth()->user()->id === $review->user_id)
                                                by <strong>you</strong>
                                            @endif
                                        </div>

                                        <div class="float-right">
                                            Posted {{ $review->created_at->diffForHumans() }}
                                        </div>

                                        <br/>
                                        <br/>

                                        <p>
                                            {{ $review->content }}
                                        </p>

                                    </div>
                                </div>
                                <div class="card-footer">
                                    <a href="{{ route('flaggedReviewStore', $review->id) }}" class="pull-right btn btn-sm btn-default">
                                        <small>Flag this review</small>
                                    </a>
                                </div>
                                @else
                                <div class="card-body">
                                    <small>This comment is currently under review.</small>
                                </div>
                                @endif
                            </div>
                        @empty
                            No reviews for this product.
                        @endforelse

                    </div>
                </div>
                <div class="card-footer">
                    @if($permissionToReview)
                        @include('layouts.errors')

                        <form action="{{ route('reviewCreate', $product->id) }}" method='POST'>
                            {{ csrf_field() }}
                            <div class="form-group">
                                <label>
                                    <select dusk="rating" name="rating" class='form-control'>
                                        <option value="">Choose a rating</option>
                                        <option value="0">0</option>
                                        <option value="1">1</option>
                                        <option value="2">2</option>
                                        <option value="3">3</option>
                                        <option value="4">4</option>
                                        <option value="5">5</option>
                                    </select>
                                </label>
                            </div>

                            <div class="form-group">
                                <textarea 
                                    dusk="content"
                                    class='form-control' 
                                    name='content' 
                                    type="text" 
                                    placeholder="Your review..."></textarea>
                            </div>

                            <div class="form-group pull-right">
                                <input 
                                    dusk="submit-btn"
                                    type="submit" 
                                    class='form-group btn btn-primary'
                                />
                            </div>
                        </form>
                    @endif
                </div>
            </div>
        </div>
    </div>
@stop
