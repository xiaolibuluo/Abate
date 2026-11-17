package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4750")
fun step4750(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
@JvmName("step4751")
fun step4751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
class Counter4752 {
    var current: Int by Delegates.observable(83) { _, _, new -> println(new) }

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
@JvmName("step4753")
fun step4753(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
fun <T> List<T>.secondOrNull4754(): T? = if (size > 1) this[1] else null

fun render4754(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry4755 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4755 = "registry-7"
    }
}
sealed class Outcome4756 {
    data class Success(val payload: String) : Outcome4756()
    data object Failure : Outcome4756()
}

fun describe4756(outcome: Outcome4756): String = when (outcome) {
    is Outcome4756.Success -> outcome.payload
    Outcome4756.Failure -> "failed"
}
