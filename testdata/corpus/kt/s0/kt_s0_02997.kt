package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24650")
fun step24650(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
object Registry24651 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24651 = "registry-47"
    }
}
sealed class Outcome24652 {
    data class Success(val payload: String) : Outcome24652()
    data object Failure : Outcome24652()
}

fun describe24652(outcome: Outcome24652): String = when (outcome) {
    is Outcome24652.Success -> outcome.payload
    Outcome24652.Failure -> "failed"
}
object Registry24653 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24653 = "registry-61"
    }
}
fun <T, R> List<T>.mapChecked24654(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome24655 {
    data class Success(val payload: String) : Outcome24655()
    data object Failure : Outcome24655()
}

fun describe24655(outcome: Outcome24655): String = when (outcome) {
    is Outcome24655.Success -> outcome.payload
    Outcome24655.Failure -> "failed"
}
class Counter24656 {
    var current: Int by Delegates.observable(82) { _, _, new -> println(new) }

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
