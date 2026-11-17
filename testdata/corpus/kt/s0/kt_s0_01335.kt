package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11010<A, B>(val first: A, val second: B)

fun combine11010(left: Int, right: Int): Pair11010<Int, Int> {
    return Pair11010(left + 59, right + 30)
}
fun <T> List<T>.secondOrNull11011(): T? = if (size > 1) this[1] else null

fun render11011(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked11012(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull11013(): T? = if (size > 1) this[1] else null

fun render11013(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter11014 {
    var current: Int by Delegates.observable(87) { _, _, new -> println(new) }

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
data class Pair11015<A, B>(val first: A, val second: B)

fun combine11015(left: Int, right: Int): Pair11015<Int, Int> {
    return Pair11015(left + 94, right + 24)
}
