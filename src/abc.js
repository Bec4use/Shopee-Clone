
  const postCustomer =async(data)=>{
    const response = await fetch('localhost:8080/customers', {
        method: 'POST', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(data) // body data type must match "Content-Type" header
      });
      return response
  }
  let x = await postCustomer(
    {
        "barcode":"B01238",
        "price":1050.9,
        "customer_id":2
    }
  )
  console.log(x);