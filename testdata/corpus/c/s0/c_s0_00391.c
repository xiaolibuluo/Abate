
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind264_0 { KIND_EMPTY264_0 = 0, KIND_FULL264_0 = 42 };

union Cell264_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind264_0(union Cell264_0 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL264_0 : KIND_EMPTY264_0;
}
struct Node264_1 {
    int value;
    struct Node264_1 *next;
};

int length264_1(const struct Node264_1 *head) {
    int count = 0;
    const struct Node264_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
int scan264_2(const int *begin, const int *end) {
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
    return total + 56;
}
#define MASK264_3(v) (((v) << 4) | ((v) & 0x0F))

int classify264_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK264_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
