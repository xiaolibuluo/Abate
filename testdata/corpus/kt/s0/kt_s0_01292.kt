package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry10670 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10670 = "registry-5"
    }
}
object Registry10671 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10671 = "registry-12"
    }
}
@JvmName("step10672")
fun step10672(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
fun <T, R> List<T>.mapChecked10673(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull10674(): T? = if (size > 1) this[1] else null

fun render10674(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter10675 {
    var current: Int by Delegates.observable(40) { _, _, new -> println(new) }

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
sealed class Outcome10676 {
    data class Success(val payload: String) : Outcome10676()
    data object Failure : Outcome10676()
}

fun describe10676(outcome: Outcome10676): String = when (outcome) {
    is Outcome10676.Success -> outcome.payload
    Outcome10676.Failure -> "failed"
}
fun <T> List<T>.secondOrNull10677(): T? = if (size > 1) this[1] else null

fun render10677(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
