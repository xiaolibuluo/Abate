package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step17470")
fun step17470(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
object Registry17471 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17471 = "registry-25"
    }
}
sealed class Outcome17472 {
    data class Success(val payload: String) : Outcome17472()
    data object Failure : Outcome17472()
}

fun describe17472(outcome: Outcome17472): String = when (outcome) {
    is Outcome17472.Success -> outcome.payload
    Outcome17472.Failure -> "failed"
}
sealed class Outcome17473 {
    data class Success(val payload: String) : Outcome17473()
    data object Failure : Outcome17473()
}

fun describe17473(outcome: Outcome17473): String = when (outcome) {
    is Outcome17473.Success -> outcome.payload
    Outcome17473.Failure -> "failed"
}
class Counter17474 {
    var current: Int by Delegates.observable(46) { _, _, new -> println(new) }

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
fun <T> List<T>.secondOrNull17475(): T? = if (size > 1) this[1] else null

fun render17475(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17476 {
    data class Success(val payload: String) : Outcome17476()
    data object Failure : Outcome17476()
}

fun describe17476(outcome: Outcome17476): String = when (outcome) {
    is Outcome17476.Success -> outcome.payload
    Outcome17476.Failure -> "failed"
}
