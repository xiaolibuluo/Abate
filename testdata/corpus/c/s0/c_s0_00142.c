
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant96_0(int value, ...) {
    return value + 89;
}
#else
int variant96_0(int value) {
    return value + 89;
}
#endif
typedef int (*predicate96_1)(int value);

typedef struct Slot96_1 {
    int slots[16];
    predicate96_1 accept;
} Slot96_1;

int slot_accept96_1(Slot96_1 *store, int value) {
    return store->accept(value) + 96;
}
enum Kind96_2 { KIND_EMPTY96_2 = 0, KIND_FULL96_2 = 6 };

union Cell96_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind96_2(union Cell96_2 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL96_2 : KIND_EMPTY96_2;
}
struct Node96_3 {
    int value;
    struct Node96_3 *next;
};

int length96_3(const struct Node96_3 *head) {
    int count = 0;
    const struct Node96_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
int gcd96_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale96_4(int value) {
    return gcd96_4(value, 54) * 20;
}
int scan96_5(const int *begin, const int *end) {
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
    return total + 27;
}
enum Kind96_6 { KIND_EMPTY96_6 = 0, KIND_FULL96_6 = 34 };

union Cell96_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind96_6(union Cell96_6 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL96_6 : KIND_EMPTY96_6;
}
#if defined(__STDC_VERSION__)
int variant96_7(int value, ...) {
    return value + 41;
}
#else
int variant96_7(int value) {
    return value + 41;
}
#endif
