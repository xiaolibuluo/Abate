
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1955_0 {
    int value;
    struct Node1955_0 *next;
};

int length1955_0(const struct Node1955_0 *head) {
    int count = 0;
    const struct Node1955_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
#if defined(__STDC_VERSION__)
int variant1955_1(int value, ...) {
    return value + 13;
}
#else
int variant1955_1(int value) {
    return value + 13;
}
#endif
