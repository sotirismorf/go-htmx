// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

type TemplAdminPage struct {
	Title       string
	CTName      string
	CanDelete   bool
	CanDownload bool
	Pagination  TemplPagination
	Columns     []string
	Rows        []TemplAdminTableRow
}

type TemplAdminTableRow struct {
	ID    string
	Cells [][]TemplAdminTableCell
}

type TemplAdminTableCell struct {
	Text       string
	IsRelation bool
	CTName     string
	ID         string
}

func AdminPage(props TemplAdminPage) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"search-results\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Breadcrumb(props.CTName, props.Title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Pagination(props.Pagination, 1).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-black sm:rounded-md border-y sm:border border-white sm:border-r overflow-x-scroll\"><table class=\"table-auto w-full border-spacing-4\"><thead class=\"text-left border-b\"><tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, row := range props.Columns {
			templ_7745c5c3_Err = TableHeader(row).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = TableHeader("").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, row := range props.Rows {
			templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				for _, cell := range row.Cells {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<td class=\"whitespace-nowrap overflow-hidden overflow-ellipsis px-4 py-2 max-w-56\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					for i, cellItem := range cell {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"hover:underline\"")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if cellItem.IsRelation {
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" href=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var3 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/admin/%s/%s", cellItem.CTName, cellItem.ID))
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						} else {
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" href=\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/admin/%s/%s", props.CTName, row.ID))
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var5 string
						templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(cellItem.Text)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/admin.templ`, Line: 55, Col: 26}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if len(cell) > 1 && i < len(cell)-1 {
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(",&nbsp")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <td class=\"whitespace-nowrap overflow-hidden overflow-ellipsis px-2 py-1 max-w-56\"><div class=\"flex items-center justify-end gap-2\"><button hx-delete=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/admin/%s/%s", props.CTName, row.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/admin.templ`, Line: 66, Col: 71}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"closest tr\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = IconDelete().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> <button hx-delete=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/admin/%s/%s", props.CTName, row.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/admin.templ`, Line: 72, Col: 71}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"closest tr\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = IconEdit("size-6").Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> <button hx-delete=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/admin/%s/%s", props.CTName, row.ID))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/admin.templ`, Line: 78, Col: 71}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"closest tr\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = IconDelete().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></div></td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = TableRow().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AdminSideNav() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"divide-y divide-white\"><h2 class=\"text-xl font-bold flex gap-2 items-center\"><span class=\"px-2 mr-4\">Material </span> <button class=\"border-l border-white p-2 hover:bg-neutral-900\" _=\"on click hide #admin-navigation-sidebar\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = IconXMark("size-6").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button></h2><ul class=\"text-lg pt-1\"><li class=\"py-0.5 px-2 hover:bg-neutral-900\"><a class=\"hover:underline \" href=\"/admin/uploads\">Uploads</a></li><li class=\"py-0.5 px-2 hover:bg-neutral-900\"><a class=\"hover:underline \" href=\"/admin/items\">Items</a></li><li class=\"py-0.5 px-2 hover:bg-neutral-900\"><a class=\"hover:underline \" href=\"/admin/authors\">Authors</a></li><li class=\"py-0.5 px-2 hover:bg-neutral-900\"><a class=\"hover:underline \" href=\"/admin/groups\">Groups</a></li></ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
