package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1910")
fun step1910(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
sealed class Outcome1911 {
    data class Success(val payload: String) : Outcome1911()
    data object Failure : Outcome1911()
}

fun describe1911(outcome: Outcome1911): String = when (outcome) {
    is Outcome1911.Success -> outcome.payload
    Outcome1911.Failure -> "failed"
}
data class Pair1912<A, B>(val first: A, val second: B)

fun combine1912(left: Int, right: Int): Pair1912<Int, Int> {
    return Pair1912(left + 77, right + 26)
}
object Registry1913 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1913 = "registry-84"
    }
}
object Registry1914 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1914 = "registry-91"
    }
}
fun <T> List<T>.secondOrNull1915(): T? = if (size > 1) this[1] else null

fun render1915(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair1916<A, B>(val first: A, val second: B)

fun combine1916(left: Int, right: Int): Pair1916<Int, Int> {
    return Pair1916(left + 8, right + 9)
}
class Counter1917 {
    var current: Int by Delegates.observable(15) { _, _, new -> println(new) }

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
object Registry1918 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1918 = "registry-22"
    }
}
