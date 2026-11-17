package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry440 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG440 = "registry-92"
    }
}
fun <T> List<T>.secondOrNull441(): T? = if (size > 1) this[1] else null

fun render441(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry442 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG442 = "registry-9"
    }
}
sealed class Outcome443 {
    data class Success(val payload: String) : Outcome443()
    data object Failure : Outcome443()
}

fun describe443(outcome: Outcome443): String = when (outcome) {
    is Outcome443.Success -> outcome.payload
    Outcome443.Failure -> "failed"
}
class Counter444 {
    var current: Int by Delegates.observable(23) { _, _, new -> println(new) }

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
@JvmName("step445")
fun step445(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
sealed class Outcome446 {
    data class Success(val payload: String) : Outcome446()
    data object Failure : Outcome446()
}

fun describe446(outcome: Outcome446): String = when (outcome) {
    is Outcome446.Success -> outcome.payload
    Outcome446.Failure -> "failed"
}
@JvmName("step447")
fun step447(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
