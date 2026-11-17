
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind470_0 { KIND_EMPTY470_0 = 0, KIND_FULL470_0 = 101 };

union Cell470_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind470_0(union Cell470_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL470_0 : KIND_EMPTY470_0;
}
struct Node470_1 {
    int value;
    struct Node470_1 *next;
};

int length470_1(const struct Node470_1 *head) {
    int count = 0;
    const struct Node470_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node470_2 {
    int value;
    struct Node470_2 *next;
};

int length470_2(const struct Node470_2 *head) {
    int count = 0;
    const struct Node470_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
enum Kind470_3 { KIND_EMPTY470_3 = 0, KIND_FULL470_3 = 25 };

union Cell470_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind470_3(union Cell470_3 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL470_3 : KIND_EMPTY470_3;
}
#if defined(__STDC_VERSION__)
int variant470_4(int value, ...) {
    return value + 32;
}
#else
int variant470_4(int value) {
    return value + 32;
}
#endif
enum Kind470_5 { KIND_EMPTY470_5 = 0, KIND_FULL470_5 = 39 };

union Cell470_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind470_5(union Cell470_5 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL470_5 : KIND_EMPTY470_5;
}
