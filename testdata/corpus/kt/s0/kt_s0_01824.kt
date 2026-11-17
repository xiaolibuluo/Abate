package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15000 {
    data class Success(val payload: String) : Outcome15000()
    data object Failure : Outcome15000()
}

fun describe15000(outcome: Outcome15000): String = when (outcome) {
    is Outcome15000.Success -> outcome.payload
    Outcome15000.Failure -> "failed"
}
@JvmName("step15001")
fun step15001(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
@JvmName("step15002")
fun step15002(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
object Registry15003 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15003 = "registry-29"
    }
}
fun <T> List<T>.secondOrNull15004(): T? = if (size > 1) this[1] else null

fun render15004(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry15005 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15005 = "registry-43"
    }
}
@JvmName("step15006")
fun step15006(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
object Registry15007 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15007 = "registry-57"
    }
}
fun <T, R> List<T>.mapChecked15008(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
