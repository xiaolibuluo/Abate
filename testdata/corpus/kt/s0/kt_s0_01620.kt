package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter13370 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
data class Pair13371<A, B>(val first: A, val second: B)

fun combine13371(left: Int, right: Int): Pair13371<Int, Int> {
    return Pair13371(left + 30, right + 46)
}
fun <T> List<T>.secondOrNull13372(): T? = if (size > 1) this[1] else null

fun render13372(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13373<A, B>(val first: A, val second: B)

fun combine13373(left: Int, right: Int): Pair13373<Int, Int> {
    return Pair13373(left + 44, right + 68)
}
fun <T, R> List<T>.mapChecked13374(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
