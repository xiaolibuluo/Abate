package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23080 {
    var current: Int by Delegates.observable(36) { _, _, new -> println(new) }

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
@JvmName("step23081")
fun step23081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
@JvmName("step23082")
fun step23082(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
fun <T> List<T>.secondOrNull23083(): T? = if (size > 1) this[1] else null

fun render23083(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23084")
fun step23084(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 64
}
