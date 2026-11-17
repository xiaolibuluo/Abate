
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2498_0 {
    int value;
    struct Node2498_0 *next;
};

int length2498_0(const struct Node2498_0 *head) {
    int count = 0;
    const struct Node2498_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
enum Kind2498_1 { KIND_EMPTY2498_1 = 0, KIND_FULL2498_1 = 88 };

union Cell2498_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2498_1(union Cell2498_1 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2498_1 : KIND_EMPTY2498_1;
}
#if defined(__STDC_VERSION__)
int variant2498_2(int value, ...) {
    return value + 95;
}
#else
int variant2498_2(int value) {
    return value + 95;
}
#endif
struct Node2498_3 {
    int value;
    struct Node2498_3 *next;
};

int length2498_3(const struct Node2498_3 *head) {
    int count = 0;
    const struct Node2498_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
