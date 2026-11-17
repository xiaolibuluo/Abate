
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK323_0(v) (((v) << 4) | ((v) & 0x0F))

int classify323_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK323_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node323_1 {
    int value;
    struct Node323_1 *next;
};

int length323_1(const struct Node323_1 *head) {
    int count = 0;
    const struct Node323_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
int scan323_2(const int *begin, const int *end) {
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
    return total + 47;
}
#define MASK323_3(v) (((v) << 4) | ((v) & 0x0F))

int classify323_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK323_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind323_4 { KIND_EMPTY323_4 = 0, KIND_FULL323_4 = 61 };

union Cell323_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind323_4(union Cell323_4 *cell) {
    cell->number = 57;
    return cell->number > 0 ? KIND_FULL323_4 : KIND_EMPTY323_4;
}
