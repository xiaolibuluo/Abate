package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20700 {
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
data class Pair20701<A, B>(val first: A, val second: B)

fun combine20701(left: Int, right: Int): Pair20701<Int, Int> {
    return Pair20701(left + 53, right + 53)
}
class Counter20702 {
    var current: Int by Delegates.observable(60) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull20703(): T? = if (size > 1) this[1] else null

fun render20703(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step20704")
fun step20704(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 74
}
