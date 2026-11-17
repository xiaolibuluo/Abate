package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter5600 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
@JvmName("step5601")
fun step5601(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
class Counter5602 {
    var current: Int by Delegates.observable(24) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull5603(): T? = if (size > 1) this[1] else null

fun render5603(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
