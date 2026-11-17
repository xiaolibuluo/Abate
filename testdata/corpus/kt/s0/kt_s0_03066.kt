package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25160 {
    data class Success(val payload: String) : Outcome25160()
    data object Failure : Outcome25160()
}

fun describe25160(outcome: Outcome25160): String = when (outcome) {
    is Outcome25160.Success -> outcome.payload
    Outcome25160.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25161(): T? = if (size > 1) this[1] else null

fun render25161(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25162<A, B>(val first: A, val second: B)

fun combine25162(left: Int, right: Int): Pair25162<Int, Int> {
    return Pair25162(left + 38, right + 14)
}
fun <T> List<T>.secondOrNull25163(): T? = if (size > 1) this[1] else null

fun render25163(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair25164<A, B>(val first: A, val second: B)

fun combine25164(left: Int, right: Int): Pair25164<Int, Int> {
    return Pair25164(left + 52, right + 36)
}
object Registry25165 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25165 = "registry-59"
    }
}
sealed class Outcome25166 {
    data class Success(val payload: String) : Outcome25166()
    data object Failure : Outcome25166()
}

fun describe25166(outcome: Outcome25166): String = when (outcome) {
    is Outcome25166.Success -> outcome.payload
    Outcome25166.Failure -> "failed"
}
