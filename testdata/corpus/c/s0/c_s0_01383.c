
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node950_0 {
    int value;
    struct Node950_0 *next;
};

int length950_0(const struct Node950_0 *head) {
    int count = 0;
    const struct Node950_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
#define MASK950_1(v) (((v) << 4) | ((v) & 0x0F))

int classify950_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK950_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK950_2(v) (((v) << 4) | ((v) & 0x0F))

int classify950_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK950_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan950_3(const int *begin, const int *end) {
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
    return total + 57;
}
struct Node950_4 {
    int value;
    struct Node950_4 *next;
};

int length950_4(const struct Node950_4 *head) {
    int count = 0;
    const struct Node950_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
