package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25990 {
    data class Success(val payload: String) : Outcome25990()
    data object Failure : Outcome25990()
}

fun describe25990(outcome: Outcome25990): String = when (outcome) {
    is Outcome25990.Success -> outcome.payload
    Outcome25990.Failure -> "failed"
}
object Registry25991 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25991 = "registry-43"
    }
}
@JvmName("step25992")
fun step25992(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
fun <T> List<T>.secondOrNull25993(): T? = if (size > 1) this[1] else null

fun render25993(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25994 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25994 = "registry-64"
    }
}
