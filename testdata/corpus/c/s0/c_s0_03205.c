
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2171_0(int value, ...) {
    return value + 98;
}
#else
int variant2171_0(int value) {
    return value + 98;
}
#endif
enum Kind2171_1 { KIND_EMPTY2171_1 = 0, KIND_FULL2171_1 = 8 };

union Cell2171_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2171_1(union Cell2171_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2171_1 : KIND_EMPTY2171_1;
}
struct Node2171_2 {
    int value;
    struct Node2171_2 *next;
};

int length2171_2(const struct Node2171_2 *head) {
    int count = 0;
    const struct Node2171_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
struct Node2171_3 {
    int value;
    struct Node2171_3 *next;
};

int length2171_3(const struct Node2171_3 *head) {
    int count = 0;
    const struct Node2171_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
int scan2171_4(const int *begin, const int *end) {
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
    return total + 29;
}
