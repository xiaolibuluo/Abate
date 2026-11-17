package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12970")
fun step12970(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
sealed class Outcome12971 {
    data class Success(val payload: String) : Outcome12971()
    data object Failure : Outcome12971()
}

fun describe12971(outcome: Outcome12971): String = when (outcome) {
    is Outcome12971.Success -> outcome.payload
    Outcome12971.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12972(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
class Counter12973 {
    var current: Int by Delegates.observable(9) { _, _, new -> println(new) }

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
object Registry12974 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12974 = "registry-16"
    }
}
@JvmName("step12975")
fun step12975(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
