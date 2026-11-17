package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull4650(): T? = if (size > 1) this[1] else null

fun render4650(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair4651<A, B>(val first: A, val second: B)

fun combine4651(left: Int, right: Int): Pair4651<Int, Int> {
    return Pair4651(left + 43, right + 42)
}
class Counter4652 {
    var current: Int by Delegates.observable(50) { _, _, new -> println(new) }

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
data class Pair4653<A, B>(val first: A, val second: B)

fun combine4653(left: Int, right: Int): Pair4653<Int, Int> {
    return Pair4653(left + 57, right + 64)
}
data class Pair4654<A, B>(val first: A, val second: B)

fun combine4654(left: Int, right: Int): Pair4654<Int, Int> {
    return Pair4654(left + 64, right + 14)
}
fun <T> List<T>.secondOrNull4655(): T? = if (size > 1) this[1] else null

fun render4655(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
