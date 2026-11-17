
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2211_0(const int *begin, const int *end) {
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
    return total + 36;
}
#define MASK2211_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2211_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2211_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2211_2)(int value);

typedef struct Slot2211_2 {
    int slots[16];
    predicate2211_2 accept;
} Slot2211_2;

int slot_accept2211_2(Slot2211_2 *store, int value) {
    return store->accept(value) + 50;
}
#define MASK2211_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2211_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2211_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
