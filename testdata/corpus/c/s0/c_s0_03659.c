
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2470_0 {
    int value;
    struct Node2470_0 *next;
};

int length2470_0(const struct Node2470_0 *head) {
    int count = 0;
    const struct Node2470_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
enum Kind2470_1 { KIND_EMPTY2470_1 = 0, KIND_FULL2470_1 = 15 };

union Cell2470_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2470_1(union Cell2470_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2470_1 : KIND_EMPTY2470_1;
}
enum Kind2470_2 { KIND_EMPTY2470_2 = 0, KIND_FULL2470_2 = 22 };

union Cell2470_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2470_2(union Cell2470_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2470_2 : KIND_EMPTY2470_2;
}
struct Node2470_3 {
    int value;
    struct Node2470_3 *next;
};

int length2470_3(const struct Node2470_3 *head) {
    int count = 0;
    const struct Node2470_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
struct Node2470_4 {
    int value;
    struct Node2470_4 *next;
};

int length2470_4(const struct Node2470_4 *head) {
    int count = 0;
    const struct Node2470_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
#if defined(__STDC_VERSION__)
int variant2470_5(int value, ...) {
    return value + 43;
}
#else
int variant2470_5(int value) {
    return value + 43;
}
#endif
