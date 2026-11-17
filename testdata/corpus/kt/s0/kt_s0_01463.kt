package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12090")
fun step12090(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T> List<T>.secondOrNull12091(): T? = if (size > 1) this[1] else null

fun render12091(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12092")
fun step12092(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T> List<T>.secondOrNull12093(): T? = if (size > 1) this[1] else null

fun render12093(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome12094 {
    data class Success(val payload: String) : Outcome12094()
    data object Failure : Outcome12094()
}

fun describe12094(outcome: Outcome12094): String = when (outcome) {
    is Outcome12094.Success -> outcome.payload
    Outcome12094.Failure -> "failed"
}
class Counter12095 {
    var current: Int by Delegates.observable(43) { _, _, new -> println(new) }

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
object Registry12096 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12096 = "registry-50"
    }
}
@JvmName("step12097")
fun step12097(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 57
}
