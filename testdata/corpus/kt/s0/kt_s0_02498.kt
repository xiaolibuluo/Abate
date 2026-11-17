package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20540")
fun step20540(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
fun <T> List<T>.secondOrNull20541(): T? = if (size > 1) this[1] else null

fun render20541(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter20542 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
