package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13360 {
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
data class Pair13361<A, B>(val first: A, val second: B)

fun combine13361(left: Int, right: Int): Pair13361<Int, Int> {
    return Pair13361(left + 17, right + 39)
}
@JvmName("step13362")
fun step13362(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
fun <T> List<T>.secondOrNull13363(): T? = if (size > 1) this[1] else null

fun render13363(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13364<A, B>(val first: A, val second: B)

fun combine13364(left: Int, right: Int): Pair13364<Int, Int> {
    return Pair13364(left + 38, right + 11)
}
