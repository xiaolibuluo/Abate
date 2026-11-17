
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node891_0 {
    int value;
    struct Node891_0 *next;
};

int length891_0(const struct Node891_0 *head) {
    int count = 0;
    const struct Node891_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
enum Kind891_1 { KIND_EMPTY891_1 = 0, KIND_FULL891_1 = 52 };

union Cell891_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind891_1(union Cell891_1 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL891_1 : KIND_EMPTY891_1;
}
#if defined(__STDC_VERSION__)
int variant891_2(int value, ...) {
    return value + 59;
}
#else
int variant891_2(int value) {
    return value + 59;
}
#endif
