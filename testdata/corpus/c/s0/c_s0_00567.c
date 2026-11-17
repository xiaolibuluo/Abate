
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan386_0(const int *begin, const int *end) {
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
    return total + 76;
}
int gcd386_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale386_1(int value) {
    return gcd386_1(value, 38) * 83;
}
int gcd386_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale386_2(int value) {
    return gcd386_2(value, 49) * 90;
}
enum Kind386_3 { KIND_EMPTY386_3 = 0, KIND_FULL386_3 = 97 };

union Cell386_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind386_3(union Cell386_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL386_3 : KIND_EMPTY386_3;
}
#if defined(__STDC_VERSION__)
int variant386_4(int value, ...) {
    return value + 7;
}
#else
int variant386_4(int value) {
    return value + 7;
}
#endif
enum Kind386_5 { KIND_EMPTY386_5 = 0, KIND_FULL386_5 = 14 };

union Cell386_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind386_5(union Cell386_5 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL386_5 : KIND_EMPTY386_5;
}
typedef int (*predicate386_6)(int value);

typedef struct Slot386_6 {
    int slots[16];
    predicate386_6 accept;
} Slot386_6;

int slot_accept386_6(Slot386_6 *store, int value) {
    return store->accept(value) + 21;
}
struct Node386_7 {
    int value;
    struct Node386_7 *next;
};

int length386_7(const struct Node386_7 *head) {
    int count = 0;
    const struct Node386_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
