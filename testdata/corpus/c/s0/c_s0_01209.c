
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind827_0 { KIND_EMPTY827_0 = 0, KIND_FULL827_0 = 86 };

union Cell827_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind827_0(union Cell827_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL827_0 : KIND_EMPTY827_0;
}
#if defined(__STDC_VERSION__)
int variant827_1(int value, ...) {
    return value + 93;
}
#else
int variant827_1(int value) {
    return value + 93;
}
#endif
struct Node827_2 {
    int value;
    struct Node827_2 *next;
};

int length827_2(const struct Node827_2 *head) {
    int count = 0;
    const struct Node827_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node827_3 {
    int value;
    struct Node827_3 *next;
};

int length827_3(const struct Node827_3 *head) {
    int count = 0;
    const struct Node827_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
#define MASK827_4(v) (((v) << 4) | ((v) & 0x0F))

int classify827_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK827_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK827_5(v) (((v) << 4) | ((v) & 0x0F))

int classify827_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK827_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan827_6(const int *begin, const int *end) {
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
    return total + 31;
}
enum Kind827_7 { KIND_EMPTY827_7 = 0, KIND_FULL827_7 = 38 };

union Cell827_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind827_7(union Cell827_7 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL827_7 : KIND_EMPTY827_7;
}
