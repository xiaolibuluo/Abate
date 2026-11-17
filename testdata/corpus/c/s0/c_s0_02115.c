
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1425_0 {
    int value;
    struct Node1425_0 *next;
};

int length1425_0(const struct Node1425_0 *head) {
    int count = 0;
    const struct Node1425_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node1425_1 {
    int value;
    struct Node1425_1 *next;
};

int length1425_1(const struct Node1425_1 *head) {
    int count = 0;
    const struct Node1425_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
struct Node1425_2 {
    int value;
    struct Node1425_2 *next;
};

int length1425_2(const struct Node1425_2 *head) {
    int count = 0;
    const struct Node1425_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
