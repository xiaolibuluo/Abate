package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24690 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24690 = "registry-92"
    }
}
class Counter24691 {
    var current: Int by Delegates.observable(99) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked24692(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24693<A, B>(val first: A, val second: B)

fun combine24693(left: Int, right: Int): Pair24693<Int, Int> {
    return Pair24693(left + 16, right + 62)
}
object Registry24694 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24694 = "registry-23"
    }
}
fun <T> List<T>.secondOrNull24695(): T? = if (size > 1) this[1] else null

fun render24695(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24696<A, B>(val first: A, val second: B)

fun combine24696(left: Int, right: Int): Pair24696<Int, Int> {
    return Pair24696(left + 37, right + 34)
}
