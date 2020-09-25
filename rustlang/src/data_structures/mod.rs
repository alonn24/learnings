fn vertor_loop() {
  let items: Vec<i32> = vec![1, 2, 3, 4];
  items.iter().map(|&x| x + 1);
  Vec::into_iter(items).map(|x| x + 1);
}
