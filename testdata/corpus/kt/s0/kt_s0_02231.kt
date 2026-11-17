package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18350")
fun step18350(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
fun <T, R> List<T>.mapChecked18351(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair18352<A, B>(val first: A, val second: B)

fun combine18352(left: Int, right: Int): Pair18352<Int, Int> {
    return Pair18352(left + 12, right + 66)
}
object Registry18353 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18353 = "registry-19"
    }
}
sealed class Outcome18354 {
    data class Success(val payload: String) : Outcome18354()
    data object Failure : Outcome18354()
}

fun describe18354(outcome: Outcome18354): String = when (outcome) {
    is Outcome18354.Success -> outcome.payload
    Outcome18354.Failure -> "failed"
}
class Counter18355 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

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
sealed class Outcome18356 {
    data class Success(val payload: String) : Outcome18356()
    data object Failure : Outcome18356()
}

fun describe18356(outcome: Outcome18356): String = when (outcome) {
    is Outcome18356.Success -> outcome.payload
    Outcome18356.Failure -> "failed"
}
