package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair1800<A, B>(val first: A, val second: B)

fun combine1800(left: Int, right: Int): Pair1800<Int, Int> {
    return Pair1800(left + 17, right + 49)
}
sealed class Outcome1801 {
    data class Success(val payload: String) : Outcome1801()
    data object Failure : Outcome1801()
}

fun describe1801(outcome: Outcome1801): String = when (outcome) {
    is Outcome1801.Success -> outcome.payload
    Outcome1801.Failure -> "failed"
}
@JvmName("step1802")
fun step1802(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
data class Pair1803<A, B>(val first: A, val second: B)

fun combine1803(left: Int, right: Int): Pair1803<Int, Int> {
    return Pair1803(left + 38, right + 21)
}
fun <T> List<T>.secondOrNull1804(): T? = if (size > 1) this[1] else null

fun render1804(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter1805 {
    var current: Int by Delegates.observable(52) { _, _, new -> println(new) }

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
object Registry1806 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1806 = "registry-59"
    }
}
