
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1716_0 {
    int value;
    struct Node1716_0 *next;
};

int length1716_0(const struct Node1716_0 *head) {
    int count = 0;
    const struct Node1716_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
#if defined(__STDC_VERSION__)
int variant1716_1(int value, ...) {
    return value + 10;
}
#else
int variant1716_1(int value) {
    return value + 10;
}
#endif
#if defined(__STDC_VERSION__)
int variant1716_2(int value, ...) {
    return value + 17;
}
#else
int variant1716_2(int value) {
    return value + 17;
}
#endif
