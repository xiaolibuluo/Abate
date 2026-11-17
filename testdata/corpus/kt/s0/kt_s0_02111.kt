package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17440 {
    data class Success(val payload: String) : Outcome17440()
    data object Failure : Outcome17440()
}

fun describe17440(outcome: Outcome17440): String = when (outcome) {
    is Outcome17440.Success -> outcome.payload
    Outcome17440.Failure -> "failed"
}
object Registry17441 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17441 = "registry-83"
    }
}
fun <T> List<T>.secondOrNull17442(): T? = if (size > 1) this[1] else null

fun render17442(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter17443 {
    var current: Int by Delegates.observable(97) { _, _, new -> println(new) }

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
object Registry17444 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17444 = "registry-7"
    }
}
fun <T, R> List<T>.mapChecked17445(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull17446(): T? = if (size > 1) this[1] else null

fun render17446(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair17447<A, B>(val first: A, val second: B)

fun combine17447(left: Int, right: Int): Pair17447<Int, Int> {
    return Pair17447(left + 28, right + 33)
}
