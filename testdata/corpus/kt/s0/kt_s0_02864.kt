package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23630 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23630 = "registry-72"
    }
}
@JvmName("step23631")
fun step23631(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 79
}
fun <T, R> List<T>.mapChecked23632(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter23633 {
    var current: Int by Delegates.observable(93) { _, _, new -> println(new) }

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
@JvmName("step23634")
fun step23634(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
sealed class Outcome23635 {
    data class Success(val payload: String) : Outcome23635()
    data object Failure : Outcome23635()
}

fun describe23635(outcome: Outcome23635): String = when (outcome) {
    is Outcome23635.Success -> outcome.payload
    Outcome23635.Failure -> "failed"
}
