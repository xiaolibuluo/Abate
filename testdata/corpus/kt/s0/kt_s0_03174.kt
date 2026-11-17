package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry26070 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26070 = "registry-43"
    }
}
fun <T> List<T>.secondOrNull26071(): T? = if (size > 1) this[1] else null

fun render26071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry26072 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26072 = "registry-57"
    }
}
sealed class Outcome26073 {
    data class Success(val payload: String) : Outcome26073()
    data object Failure : Outcome26073()
}

fun describe26073(outcome: Outcome26073): String = when (outcome) {
    is Outcome26073.Success -> outcome.payload
    Outcome26073.Failure -> "failed"
}
fun <T> List<T>.secondOrNull26074(): T? = if (size > 1) this[1] else null

fun render26074(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry26075 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG26075 = "registry-78"
    }
}
