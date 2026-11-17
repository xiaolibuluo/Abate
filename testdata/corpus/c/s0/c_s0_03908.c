
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2647_0(int value, ...) {
    return value + 78;
}
#else
int variant2647_0(int value) {
    return value + 78;
}
#endif
#if defined(__STDC_VERSION__)
int variant2647_1(int value, ...) {
    return value + 85;
}
#else
int variant2647_1(int value) {
    return value + 85;
}
#endif
struct Node2647_2 {
    int value;
    struct Node2647_2 *next;
};

int length2647_2(const struct Node2647_2 *head) {
    int count = 0;
    const struct Node2647_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
#if defined(__STDC_VERSION__)
int variant2647_3(int value, ...) {
    return value + 99;
}
#else
int variant2647_3(int value) {
    return value + 99;
}
#endif
#if defined(__STDC_VERSION__)
int variant2647_4(int value, ...) {
    return value + 9;
}
#else
int variant2647_4(int value) {
    return value + 9;
}
#endif
