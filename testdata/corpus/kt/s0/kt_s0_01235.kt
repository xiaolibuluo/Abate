package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter10190 {
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
@JvmName("step10191")
fun step10191(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 67
}
fun <T> List<T>.secondOrNull10192(): T? = if (size > 1) this[1] else null

fun render10192(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair10193<A, B>(val first: A, val second: B)

fun combine10193(left: Int, right: Int): Pair10193<Int, Int> {
    return Pair10193(left + 81, right + 38)
}
data class Pair10194<A, B>(val first: A, val second: B)

fun combine10194(left: Int, right: Int): Pair10194<Int, Int> {
    return Pair10194(left + 88, right + 49)
}
@JvmName("step10195")
fun step10195(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
data class Pair10196<A, B>(val first: A, val second: B)

fun combine10196(left: Int, right: Int): Pair10196<Int, Int> {
    return Pair10196(left + 5, right + 10)
}
@JvmName("step10197")
fun step10197(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
data class Pair10198<A, B>(val first: A, val second: B)

fun combine10198(left: Int, right: Int): Pair10198<Int, Int> {
    return Pair10198(left + 19, right + 32)
}
class Counter10199 {
    var current: Int by Delegates.observable(26) { _, _, new -> println(new) }

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
