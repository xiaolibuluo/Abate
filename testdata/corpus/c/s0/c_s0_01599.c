
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1079_0(int value, ...) {
    return value + 64;
}
#else
int variant1079_0(int value) {
    return value + 64;
}
#endif
#if defined(__STDC_VERSION__)
int variant1079_1(int value, ...) {
    return value + 71;
}
#else
int variant1079_1(int value) {
    return value + 71;
}
#endif
struct Node1079_2 {
    int value;
    struct Node1079_2 *next;
};

int length1079_2(const struct Node1079_2 *head) {
    int count = 0;
    const struct Node1079_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
