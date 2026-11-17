package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step10620")
fun step10620(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 37
}
object Registry10621 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10621 = "registry-44"
    }
}
fun <T, R> List<T>.mapChecked10622(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry10623 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10623 = "registry-58"
    }
}
sealed class Outcome10624 {
    data class Success(val payload: String) : Outcome10624()
    data object Failure : Outcome10624()
}

fun describe10624(outcome: Outcome10624): String = when (outcome) {
    is Outcome10624.Success -> outcome.payload
    Outcome10624.Failure -> "failed"
}
object Registry10625 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10625 = "registry-72"
    }
}
class Counter10626 {
    var current: Int by Delegates.observable(79) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked10627(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
