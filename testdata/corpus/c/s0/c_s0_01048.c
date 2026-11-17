
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK718_0(v) (((v) << 4) | ((v) & 0x0F))

int classify718_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK718_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK718_1(v) (((v) << 4) | ((v) & 0x0F))

int classify718_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK718_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind718_2 { KIND_EMPTY718_2 = 0, KIND_FULL718_2 = 41 };

union Cell718_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind718_2(union Cell718_2 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL718_2 : KIND_EMPTY718_2;
}
#if defined(__STDC_VERSION__)
int variant718_3(int value, ...) {
    return value + 48;
}
#else
int variant718_3(int value) {
    return value + 48;
}
#endif
#define MASK718_4(v) (((v) << 4) | ((v) & 0x0F))

int classify718_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK718_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan718_5(const int *begin, const int *end) {
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
    return total + 62;
}
struct Node718_6 {
    int value;
    struct Node718_6 *next;
};

int length718_6(const struct Node718_6 *head) {
    int count = 0;
    const struct Node718_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
