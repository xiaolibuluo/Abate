package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter16260 {
    var current: Int by Delegates.observable(94) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step16261")
fun step16261(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 101
}
fun <T> List<T>.secondOrNull16262(): T? = if (size > 1) this[1] else null

fun render16262(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16263")
fun step16263(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
