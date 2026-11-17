package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter3890 {
    var current: Int by Delegates.observable(18) { _, _, new -> println(new) }

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
sealed class Outcome3891 {
    data class Success(val payload: String) : Outcome3891()
    data object Failure : Outcome3891()
}

fun describe3891(outcome: Outcome3891): String = when (outcome) {
    is Outcome3891.Success -> outcome.payload
    Outcome3891.Failure -> "failed"
}
sealed class Outcome3892 {
    data class Success(val payload: String) : Outcome3892()
    data object Failure : Outcome3892()
}

fun describe3892(outcome: Outcome3892): String = when (outcome) {
    is Outcome3892.Success -> outcome.payload
    Outcome3892.Failure -> "failed"
}
class Counter3893 {
    var current: Int by Delegates.observable(39) { _, _, new -> println(new) }

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
object Registry3894 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3894 = "registry-46"
    }
}
fun <T, R> List<T>.mapChecked3895(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step3896")
fun step3896(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 60
}
fun <T, R> List<T>.mapChecked3897(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
