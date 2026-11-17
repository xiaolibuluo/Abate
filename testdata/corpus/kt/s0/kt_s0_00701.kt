package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5760(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5761 {
    data class Success(val payload: String) : Outcome5761()
    data object Failure : Outcome5761()
}

fun describe5761(outcome: Outcome5761): String = when (outcome) {
    is Outcome5761.Success -> outcome.payload
    Outcome5761.Failure -> "failed"
}
sealed class Outcome5762 {
    data class Success(val payload: String) : Outcome5762()
    data object Failure : Outcome5762()
}

fun describe5762(outcome: Outcome5762): String = when (outcome) {
    is Outcome5762.Success -> outcome.payload
    Outcome5762.Failure -> "failed"
}
@JvmName("step5763")
fun step5763(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 45
}
class Counter5764 {
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
class Counter5765 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

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
object Registry5766 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5766 = "registry-66"
    }
}
