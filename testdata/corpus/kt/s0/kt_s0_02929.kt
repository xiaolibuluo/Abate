package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24160(): T? = if (size > 1) this[1] else null

fun render24160(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked24161(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24162 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24162 = "registry-96"
    }
}
sealed class Outcome24163 {
    data class Success(val payload: String) : Outcome24163()
    data object Failure : Outcome24163()
}

fun describe24163(outcome: Outcome24163): String = when (outcome) {
    is Outcome24163.Success -> outcome.payload
    Outcome24163.Failure -> "failed"
}
data class Pair24164<A, B>(val first: A, val second: B)

fun combine24164(left: Int, right: Int): Pair24164<Int, Int> {
    return Pair24164(left + 13, right + 68)
}
@JvmName("step24165")
fun step24165(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
class Counter24166 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
