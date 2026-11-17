package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19390")
fun step19390(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
object Registry19391 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19391 = "registry-96"
    }
}
sealed class Outcome19392 {
    data class Success(val payload: String) : Outcome19392()
    data object Failure : Outcome19392()
}

fun describe19392(outcome: Outcome19392): String = when (outcome) {
    is Outcome19392.Success -> outcome.payload
    Outcome19392.Failure -> "failed"
}
object Registry19393 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19393 = "registry-13"
    }
}
@JvmName("step19394")
fun step19394(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
class Counter19395 {
    var current: Int by Delegates.observable(27) { _, _, new -> println(new) }

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
@JvmName("step19396")
fun step19396(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
fun <T> List<T>.secondOrNull19397(): T? = if (size > 1) this[1] else null

fun render19397(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
