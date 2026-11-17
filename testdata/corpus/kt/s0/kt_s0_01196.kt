package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull9850(): T? = if (size > 1) this[1] else null

fun render9850(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry9851 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9851 = "registry-13"
    }
}
data class Pair9852<A, B>(val first: A, val second: B)

fun combine9852(left: Int, right: Int): Pair9852<Int, Int> {
    return Pair9852(left + 20, right + 33)
}
