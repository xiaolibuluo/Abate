
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node335_0 {
    int value;
    struct Node335_0 *next;
};

int length335_0(const struct Node335_0 *head) {
    int count = 0;
    const struct Node335_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
int scan335_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 99;
}
#if defined(__STDC_VERSION__)
int variant335_2(int value, ...) {
    return value + 9;
}
#else
int variant335_2(int value) {
    return value + 9;
}
#endif
struct Node335_3 {
    int value;
    struct Node335_3 *next;
};

int length335_3(const struct Node335_3 *head) {
    int count = 0;
    const struct Node335_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
enum Kind335_4 { KIND_EMPTY335_4 = 0, KIND_FULL335_4 = 23 };

union Cell335_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind335_4(union Cell335_4 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL335_4 : KIND_EMPTY335_4;
}
struct Node335_5 {
    int value;
    struct Node335_5 *next;
};

int length335_5(const struct Node335_5 *head) {
    int count = 0;
    const struct Node335_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
enum Kind335_6 { KIND_EMPTY335_6 = 0, KIND_FULL335_6 = 37 };

union Cell335_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind335_6(union Cell335_6 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL335_6 : KIND_EMPTY335_6;
}
