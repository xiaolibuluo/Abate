package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter23760 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
sealed class Outcome23761 {
    data class Success(val payload: String) : Outcome23761()
    data object Failure : Outcome23761()
}

fun describe23761(outcome: Outcome23761): String = when (outcome) {
    is Outcome23761.Success -> outcome.payload
    Outcome23761.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked23762(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23763 {
    data class Success(val payload: String) : Outcome23763()
    data object Failure : Outcome23763()
}

fun describe23763(outcome: Outcome23763): String = when (outcome) {
    is Outcome23763.Success -> outcome.payload
    Outcome23763.Failure -> "failed"
}
object Registry23764 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23764 = "registry-75"
    }
}
@JvmName("step23765")
fun step23765(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
class Counter23766 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
