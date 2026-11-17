package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12070 {
    data class Success(val payload: String) : Outcome12070()
    data object Failure : Outcome12070()
}

fun describe12070(outcome: Outcome12070): String = when (outcome) {
    is Outcome12070.Success -> outcome.payload
    Outcome12070.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12071(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12072 {
    data class Success(val payload: String) : Outcome12072()
    data object Failure : Outcome12072()
}

fun describe12072(outcome: Outcome12072): String = when (outcome) {
    is Outcome12072.Success -> outcome.payload
    Outcome12072.Failure -> "failed"
}
object Registry12073 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12073 = "registry-100"
    }
}
@JvmName("step12074")
fun step12074(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T> List<T>.secondOrNull12075(): T? = if (size > 1) this[1] else null

fun render12075(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12076(): T? = if (size > 1) this[1] else null

fun render12076(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter12077 {
    var current: Int by Delegates.observable(31) { _, _, new -> println(new) }

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
