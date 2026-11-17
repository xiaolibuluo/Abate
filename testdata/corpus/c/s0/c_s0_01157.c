
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK795_0(v) (((v) << 4) | ((v) & 0x0F))

int classify795_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK795_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK795_1(v) (((v) << 4) | ((v) & 0x0F))

int classify795_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK795_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node795_2 {
    int value;
    struct Node795_2 *next;
};

int length795_2(const struct Node795_2 *head) {
    int count = 0;
    const struct Node795_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#define MASK795_3(v) (((v) << 4) | ((v) & 0x0F))

int classify795_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK795_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
