package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4360")
fun step4360(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
sealed class Outcome4361 {
    data class Success(val payload: String) : Outcome4361()
    data object Failure : Outcome4361()
}

fun describe4361(outcome: Outcome4361): String = when (outcome) {
    is Outcome4361.Success -> outcome.payload
    Outcome4361.Failure -> "failed"
}
@JvmName("step4362")
fun step4362(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
object Registry4363 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4363 = "registry-68"
    }
}
@JvmName("step4364")
fun step4364(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
fun <T> List<T>.secondOrNull4365(): T? = if (size > 1) this[1] else null

fun render4365(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter4366 {
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
object Registry4367 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4367 = "registry-96"
    }
}
fun <T, R> List<T>.mapChecked4368(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
