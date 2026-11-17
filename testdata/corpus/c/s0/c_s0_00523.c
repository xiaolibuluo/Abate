
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant361_0(int value, ...) {
    return value + 42;
}
#else
int variant361_0(int value) {
    return value + 42;
}
#endif
enum Kind361_1 { KIND_EMPTY361_1 = 0, KIND_FULL361_1 = 49 };

union Cell361_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind361_1(union Cell361_1 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL361_1 : KIND_EMPTY361_1;
}
struct Node361_2 {
    int value;
    struct Node361_2 *next;
};

int length361_2(const struct Node361_2 *head) {
    int count = 0;
    const struct Node361_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
struct Node361_3 {
    int value;
    struct Node361_3 *next;
};

int length361_3(const struct Node361_3 *head) {
    int count = 0;
    const struct Node361_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
